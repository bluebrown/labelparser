package labelparser

import (
	"encoding/json"
	"testing"
)

func TestRun(t *testing.T) {
	labels := map[string]string{
		"com.docker.stack.namespace":    "my-stack",
		"com.docker.swarm.node.id":      "y7t70zw5vmylsbo5qhmc56fgb",
		"com.docker.swarm.service.id":   "dy4kw0r22a4wqunp53sb7brqs",
		"com.docker.swarm.service.name": "my-stack_manager",
		"com.docker.swarm.task":         "",
		"com.docker.swarm.task.id":      "ymx9j64986pnkmqii951a98nj",
		"com.docker.swarm.task.name":    "my-stack_manager.1.ymx9j64986pnkmqii951a98nj",
	}

	conf := map[string]interface{}{}
	Parse(labels, &conf)

	b, err := json.Marshal(conf)

	if err != nil {
		t.Fatal(err)
	}

	op := string(b)

	if op != `{"com":{"docker":{"stack":{"namespace":"my-stack"},"swarm":{"node":{"id":"y7t70zw5vmylsbo5qhmc56fgb"},"service":{"id":"dy4kw0r22a4wqunp53sb7brqs","name":"my-stack_manager"},"task":{"id":"ymx9j64986pnkmqii951a98nj","name":"my-stack_manager.1.ymx9j64986pnkmqii951a98nj","task":""}}}}}` {
		t.Fatal("output not correct")
	}

}
