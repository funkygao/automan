<?php

namespace Services;

class example {

    /**
     * @InputTag(jsonf='json/hello.i.json')
     * @OutputTag(jsonf='json/hello.o.json')
     * @param array $param
     * @return array
     */
    public function helloWorld($param) {
        return array(
            'uid' => 5,
            'status' => 'ok',
        );
    }
}