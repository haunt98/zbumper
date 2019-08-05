package main

import "os/exec"

func buildAndDeployDocker(b *Bump) (string, error) {
	/*
	image_tag=registry-gitlab.zalopay.vn/apps/zpm/backend/$1/$2:$3$now

	docker build -t $image_tag -f $1/docker/Dockerfile_$2 . && docker login registry-gitlab.zalopay.vn && docker push $image_tag
	 */
	//imageTagLocation
	//imageTag :=
	//output, err := exec.Command(
	//	"docker", "build", "-t", "$image_tag", "-f", "$1/docker/Dockerfile_$2", ".", "&&",
	//	"docker", "login", "registry-gitlab.zalopay.vn", "&&",
	//	"docker", "push", "$image_tag").
	//	CombinedOutput()

}
