<?php

namespace Service;

class example {

    /**
     * @InputTag()
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