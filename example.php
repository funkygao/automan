<?php

namespace Services;

class example {

    /**
     * @InputTag(value='hello')
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