# goutils

tiny go utilities, mainly script-like programs

## jpgScaler

Tiny Cli picture finder software written in Go. Collects all pics from rambling folders into a specified output dir

## Quick Run

Clone source code from GitHub, then build. Will be released to DockerHub after testing.

```bash
git clone https://github.com/mizumoto-cn/goutils.git
```

```bash
go get -u && go mod tidy
go build
```

then run (on Windows)

```bash
.\goutils.exe -i <input_folder_path> -o <output_target_path>
```

or Linux/Unix-like

```bash
.\goutils -i <input_folder_path> -o <output_target_path>
```

## Licensing

License: [MGPL v1.2](/License/Mizumoto%20General%20Public%20License%20v1.2.md)Basically a Mozilla 2.0 public license, but with extra restrictions:

By using any part of this project, you are deemed to have fully understanding and acceptance of the following terms：

1. You must conspicuously display, without modification, this License and the notice on each redistributed or derivative copy of the License Covered Work.
2. Any non-independent developers companies/groups/legal entities or other organizations should ensure that employees are not oppressed or exploited, and that employees can always receive a reasonable salary for their legal working hours.
3. Any independent or non-independent developers/companies/groups/legal entities or other organizations, shall ensure that it has a clear conscience, including and not limited to **opposition to any form of Nazi or Neo-Nazism organization(s)**.

Otherwise these Individuals / Companies / Groups / Legal-entities **will not have the right to copy / modify / redistribute any code / file / algorithm** governed by MGPL v1.2.