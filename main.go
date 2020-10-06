package main

import (
	"context"
	"fmt"

	registryutil "github.com/jmrodri/debug-bundle-image/pkg/registry"
)

func main() {
	bundleImage := "docker.io/jmrodri/demo-zeus-operator-bundle:0.0.1"
	fmt.Printf("Extracting %v\n", bundleImage)
	bundlePath, err := registryutil.ExtractBundleImage(context.TODO(), nil, bundleImage, false)
	if err != nil {
		fmt.Printf("pull bundle image: %v", err)
	}
	fmt.Printf("Bundle path: %v\n", bundlePath)
}
