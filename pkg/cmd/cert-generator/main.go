/*
Copyright 2023 The KusionStack Authors.
Copyright 2021 The Kruise Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"

	"k8s.io/klog/v2"

	"github.com/KusionStack/ctrlmesh/pkg/webhook/util/writer"
)

var (
	dnsName = flag.String("dns-name", "127.0.0.1", "Common name (IP or Hostname) for certs generation.")
	destDir = flag.String("dir", "/tmp/certs", "The directory to generate certs.")
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()

	certWriter, err := writer.NewFSCertWriter(writer.FSCertWriterOptions{Path: *destDir})
	if err != nil {
		klog.Fatalf("Failed to new cert writer: %v", err)
	}

	_, _, err = certWriter.EnsureCert(*dnsName)
	if err != nil {
		klog.Fatalf("Failed to generate certs into %s: %v", *dnsName, err)
	}
}
