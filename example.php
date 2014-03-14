<?php

namespace Services;

class example {

    /**
     * @InputTag(jsonf='hello.json')
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