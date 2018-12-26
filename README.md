## maul

Provides computer system monitoring, fault notification and
interim remediation.

Final remediation is handled by process improvements

- application iteration
- automated scaling
- process relocation for hardware faults

iterative automation

embrace the evolutionary process - univeral law

replace toil and waste with learning and creativity

dry kiss

dot-dash


- agate - prometheus alert distribution service

## validate

The following binaries need to exist in $GOPATH/bin

- prometheus
- alertmanager
- node_exporter
- postgres_exporter

```
go get github.com/pahoughton/maul
cd $GOPATH/src/github.com/pahoughton/maul
go get
go build
vagrant up
```

## features


- run ansible playbooks to remediate alerts
- generate tickets for alerts

## install

put 'maul' wherever you want to

## usage

```
maul --laddr :5001
curl localhost:5001/metrics
```

## contribute

[Github pahoughton/maul](https://github.com/pahoughton/maul)

## licenses

2018-12-05 (cc) <paul4hough@gmail.com>


[![LICENSE](http://i.creativecommons.org/l/by/3.0/88x31.png)](http://creativecommons.org/licenses/by/3.0/)
