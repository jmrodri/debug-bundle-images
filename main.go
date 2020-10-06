package main

import (
	"context"
	"fmt"

	registryutil "github.com/operator-framework/operator-sdk/internal/registry"
)

func main() {
	bundleImage := "docker.io/jmrodri/demo-zeus-operator-bundle:0.0.1"
	bundlePath, err := registryutil.ExtractBundleImage(context.TODO(), nil, bundleImage, false)
	if err != nil {
		return nil, nil, fmt.Errorf("pull bundle image: %v", err)
	}
	fmt.Printf("Bundle path: %v\n", bundlePath)

}
