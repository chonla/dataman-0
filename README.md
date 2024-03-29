# dataman

Random data generator.

## Usage

```
dataman gen <configuration-file.yml>
```

### Other Usages

#### Update internal datasets

dataman provides internal datasets for data generation. You can use the following command to download/update internal datasets.

```
dataman update
```

#### Show internal datasets stats

```
dataman info
```

#### Show app version

```
dataman version
```

#### Show help page

```
dataman help
```

## Configuration File (.yml)

```yml
datasets: <dataset-path>
export:
    output:
        count: <number-of-row-to-be-exported>
        to: <export-dialect>
    columns:
        - name: <field_name>
          value: <random_data>
          type: <data_type>
        ...
```

## Dataset

Data in dataset is just line-separated text in text file with `.txt` extension.

### Dataset Example

```txt
Aaron
Adam
Aidan
```

## Export Dialect

### Supported Output

#### Console

Use `console:<format>` to print random data out to console.

#### File

Use `file:<filename-with-extension>` to export random data to file. File format will be detected from filename extension.

### Supported Format/Extension

#### CSV

Use `csv` for comma separated value.

#### TSV

Use `tsv` for tab separated value.

#### JSON

Use `json` for JSON.

#### SQL

Use `sql` for SQL insert statement.

## Random Data

Random data are surrounded by `${}`. The value inside `${}` is so-called `variable`. Variable, by default, represents dataset name. The random data will be selected from that dataset.

Value other than variable will be exported as it is.

### Session Variable

Session variable is prefixed by `session`.

#### Supported Session Variables

* `${session.index}` refers to row number, equivalent to `${system.seqNumber:1}`.

### System Variable

System variable is prefixed by `system`.

#### Supported System Variables

* `${system.int:min:max}` refers to a random number between given min and max.
* `${system.decimal:min:max:<precision>}` refers to a random decimal between given min and max with given precision. Default `precision` is `5`.
* `${system.date:layout:min:max}` refers to a random date between given min and max with given layout.
* `${system.seqNumber:<start_from>}` refers to row number with specified `start_form` number. Default `start_form` is `1`.
* `${system.seqDate:layout:<start_from>}` refers to date sequence with given layout and specified `start_form` date. Default `start_form` is execution time.

## Exported Data Type

* `string` or `varchar` or `text` or omitted
* `int` or `integer`
* `number` or `decimal`
* `timestamp`

## License

[MIT](LICENSE)