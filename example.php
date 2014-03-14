<?php

namespace Services;

class example {

    /**
     * @InputTag(json='@json/hello.i.json')
     * @OutputTag(json='@json/hello.o.json')
     * @param array $param
     * @return array
     */
    public function helloWorld($param) {
        return array(
            'uid' => 5,
            'status' => 'ok',
        );
    }

    /**
     * @InputTag(json='{"uid":3}')
     * @param array $param
     */
    public function foo($param) {

    }
}