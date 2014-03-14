<?php
/**
 * Automat
 */
require_once 'bootstrap.php';

$src = '/Users/gaopeng/fun/dragon-server-code/v2/classes/Services/AccountService.php';
require_once $src;

function parseFile($filename) {
    require_once $filename;
    $clsname = 'Services\\' . substr($filename, 0, -strlen('.php'));
    echo $clsname;

    $reflection = new ReflectionClass($clsname);
    foreach ($reflection->getMethods() as $method) {
        $p = new ReflectionAnnotatedMethod($method->class, $method->name);
        $inputTag = $p->getAnnotation("InputTag");
        if (!$inputTag) {
            continue;
        }

        print_r($inputTag);
        continue;

        print_r($p->getAnnotations());
        foreach ($p->getAllAnnotations() as $a) {
            print_r($a);

        }
        //echo $p->getDocComment();
    }

}


