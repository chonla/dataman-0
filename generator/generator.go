package generator

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
)

type ResolverFunc func(args []string, sessionVars map[string]string) string

type Generator struct {
	config   *config
	datasets map[string][]string
	systems  map[string]ResolverFunc
}

func New(configFile string) (*Generator, error) {
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return nil, err
	}

	var conf = config{}

	configData, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(configData, &conf)
	if err != nil {
		return nil, err
	}

	generator := &Generator{
		config:   &conf,
		datasets: map[string][]string{},
		systems: map[string]ResolverFunc{
			"system.date":      DateResolver,
			"system.int":       IntResolver,
			"system.decimal":   DecimalResolver,
			"system.seqNumber": SeqNumberResolver,
			"system.seqDate":   SeqDateResolver,
		},
	}

	internalDatasetsPath, _ := homedir.Expand("~/.dataman/datasets")

	err = generator.LoadDatasets(internalDatasetsPath)
	if err != nil {
		return nil, err
	}

	if conf.Datasets != "" {
		err = generator.LoadDatasets(conf.Datasets)
		if err != nil {
			return nil, err
		}
	}

	return generator, nil
}

func (g *Generator) LoadDatasets(path string) error {
	files, err := filepath.Glob(fmt.Sprintf("%s/*.txt", path))
	if err != nil {
		return err
	}

	for _, file := range files {
		_, fileName := filepath.Split(file)
		datasetName := fileName[:len(fileName)-len(filepath.Ext(fileName))]
		g.LoadDataset(datasetName, file)
	}

	return nil
}

func (g *Generator) LoadDataset(name, path string) error {
	dataByte, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	data := strings.Split(string(dataByte), "\n")
	g.datasets[name] = data
	return nil
}

func (g *Generator) Generate() error {
	records := []map[string]interface{}{}
	for i := 0; i < g.config.Export.Output.Count; i++ {
		row := g.CreateRow(i)
		records = append(records, row)
	}

	output := g.config.Export.Output.To
	outputDialect := strings.Split(output, ":")
	writerType := outputDialect[0]
	outputFormat := outputDialect[1]

	headers := []string{}
	for _, colValue := range g.config.Export.Columns {
		headers = append(headers, colValue.Name)
	}

	w, err := NewWriter(writerType, outputFormat)
	if err != nil {
		return err
	}
	return w.Write(headers, records)
}

func (g *Generator) CreateRow(index int) map[string]interface{} {
	sessionVars := map[string]string{
		"session.index": fmt.Sprintf("%d", index+1),
	}
	output := map[string]interface{}{}
	for _, colValue := range g.config.Export.Columns {
		if g.containsVariable(colValue.Value) {
			output[colValue.Name] = g.renderData(colValue.Value, colValue.Type, sessionVars)
		} else {
			output[colValue.Name] = g.parseData(colValue.Value, colValue.Type)
		}
	}
	return output
}

func (g *Generator) containsVariable(val string) bool {
	re := regexp.MustCompile("\\$\\{[^}]+\\}")
	return re.Match([]byte(val))
}

func (g *Generator) tryParseCallable(val string) (ResolverFunc, []string, bool) {
	result := strings.SplitN(val, ":", 2)

	if resolver, ok := g.systems[result[0]]; ok {
		if len(result) > 1 {
			return resolver, argsSplit(result[1]), true
		}
		return resolver, []string{}, true
	}
	return nil, nil, false
}

func (g *Generator) getVariables(template string) []string {
	re := regexp.MustCompile("\\$\\{([^}]+)\\}")
	m := re.FindAllStringSubmatch(template, -1)
	if len(m) > 0 && len(m[0]) == 2 {
		matched := []string{}
		for i, ml := 0, len(m); i < ml; i++ {
			matched = append(matched, m[i][1])
		}
		return matched
	}
	return []string{}
}

func (g *Generator) renderData(template, returnType string, sessionVars map[string]string) interface{} {
	vars := g.getVariables(template)
	result := template
	rand.Seed(time.Now().UnixNano())

	for _, variable := range vars {
		replaceableVariableName := fmt.Sprintf("${%s}", variable)
		if variableValue, ok := sessionVars[variable]; ok {
			result = strings.ReplaceAll(result, replaceableVariableName, variableValue)
		} else {
			if callerFunc, callerArgs, ok := g.tryParseCallable(variable); ok {
				callResult := callerFunc(callerArgs, sessionVars)
				result = strings.ReplaceAll(result, replaceableVariableName, callResult)
			} else {
				if _, ok := g.datasets[variable]; ok {
					randomIndex := rand.Intn(len(g.datasets[variable]))
					pick := g.datasets[variable][randomIndex]
					result = strings.ReplaceAll(result, replaceableVariableName, pick)
				}
			}
		}
	}

	return g.parseData(result, returnType)
}

func (g *Generator) parseData(value, returnType string) interface{} {
	if returnType == "" || returnType == "varchar" || returnType == "string" || returnType == "text" {
		return value
	}
	if returnType == "int" || returnType == "integer" {
		intResult, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return value
		}
		return intResult
	}
	if returnType == "number" || returnType == "decimal" {
		floatResult, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return value
		}
		return floatResult
	}
	return value
}
