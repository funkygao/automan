<?php

namespace Services;

class example {

    /**
     * @In('@json/hello.i.json')
     * @Out('@json/hello.o.json')
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
     * @In('{"uid":3}')
     * @param array $param
     */
    public function foo($param) {

    }
}
