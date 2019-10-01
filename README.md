# Excel Converter

## 配置文件

配置文件放在可执行文件同层目录下的`config.yml`文件中，其中：

- `source.dir`为需要转换文件存放的目录，不支持子目录；
- `destination.dir`为转换后文件存放的目录；
- `template.file`为模板文件的路径；
- `server.port`为`server`模式的端口。

## 模板文件

其中：

- `file`为模板格式Excel文件；
- `sheets`为源文件中需要转换的表：
  - `src.start`为源文件中的起始行号；
  - `src.end`为源文件中的结束行号；
  - `dst.sheet`为目标文件的表名；
  - `dst.start`为目标文件的起始行号；
  - `mapping`为源文件列到目标文件列的映射，如源文件中的B列需要转换到目标文件的D列，则映射为2->4。

## cli mode

命令行模式，支持批处理。

```shell script
Usage of cli: converter cli [OPTIONS]
  -c string
        configuration file (default "config.yml") 自定义配置文件的路径
  -d string
        destination directory 自定义转换后文件存放的目录（覆盖配置文件）
  -s string
        source directory 自定义需要转换文件存放的目录（覆盖配置文件）
  -t string
        template file 自定义模板文件的路径（覆盖配置文件）
```

## server mode

服务模式，支持批处理。

```shell script
Usage of server: ./converter server [OPTIONS]
  -c string
        configuration file (default "config.yml") 自定义配置文件的路径
```

通过Restfull文件：

Method：POST

BODY："files": [file list]