FROM golang:latest as BUILDER

# build binary
RUN mkdir -p /go/src/openeuler/go-py

COPY . /go/src/openeuler/go-py

RUN go env -w GOPROXY=https://goproxy.cn,direct

RUN cd /go/src/openeuler/go-py && go mod tidy && CGO_ENABLED=1 go build -v -o ./go-py main.go

# copy binary config and utils
FROM openeuler/openeuler:21.03

RUN yum update && yum install -y python3 && yum install -y python3-pip

RUN mkdir -p /opt/app/go-py/py

COPY ./py /opt/app/go-py/py

RUN chmod 755 -R /opt/app/go-py/py

ENV EVALUATE /opt/app/go-py/py/evaluate.py
ENV CALCULATE /opt/app/go-py/py/calculate_fid.py
ENV UPLOAD /opt/app/go-py/py/

RUN pip install esdk-obs-python --trusted-host pypi.org
RUN pip install -r /opt/app/go-py/py/requirements.txt

COPY --from=BUILDER /go/src/openeuler/go-py/go-py /opt/app/go-py

WORKDIR /opt/app/go-py/

ENTRYPOINT ["/opt/app/go-py/go-py"]