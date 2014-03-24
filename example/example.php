<?php

namespace Services;

class example {

    /**
     *
     * @Spec('http://wiki.mycorp.com/display/DG/test+auto+wiki')
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
     *
     * @In('{"uid":3}')
     * @Out('{"status":"ok","time":139343434,"msg":"good"}')
     * @param array $param
     */
    public function foo($param) {

    }
}
