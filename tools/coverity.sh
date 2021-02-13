#!/bin/zsh

# quick go test and output coverity html report
# the first arg is the module name under directory biz, such as "article" or "album"

output_dir=$(dirname $0)"/../build/"
module_name=$1
html_report_path=${output_dir}"cover.html"
data_report_path=${output_dir}"cover.out"

if [ ! $1 ]; then
  echo "need one argument, please input the module name under directory biz which shall be tested and reported."
  exit
fi

if [ -f ${html_report_path} ]; then
  rm ${html_report_path}
fi

go test ./biz/${module_name} --cover -v --coverprofile ${data_report_path}
go tool cover -html=${data_report_path} -o ${html_report_path}

if [ -f ${data_report_path} ]; then
  rm ${data_report_path}
fi

open ${html_report_path}
