<?php
/**
 * Automatically generate api doc according our customized phpdoc tags.
 *
 * Currently supported tags:
 * InputTag
 * OutputTag
 */
require_once 'bootstrap.php';

function parsePhpFile($filename) {
    require_once $filename;
    $basecls = substr($filename, 0, -strlen('.php'));
    $clsname = 'Services\\' . $basecls;
    echo "Parsing class $clsname...\n";

    $reflection = new ReflectionClass($clsname);
    foreach ($reflection->getMethods() as $method) {
        $p = new ReflectionAnnotatedMethod($method->class, $method->name);
        $url = '/api/?class=' . $basecls . '&method=' . $method->name . '&params=';
        echo "URL: $url\n";

        foreach (array('InputTag', 'OutputTag') as $tagName) {
            $tag = $p->getAnnotation($tagName);
            if (!$tag) {
                continue;
            }



            print_r($tag);
        }
    }

}

function showHelp() {
    echo "Usage: " . $argv[0] . " filename [filename]\n";
    //exit(0);
}

function main() {
    ini_set('register_argc_argv', true);
    if (count($argv) == 0) {
        showHelp();
    }

    foreach ($argv as $filename) {
        parsePhpFile($filename);
    }

    parsePhpFile('example.php');
}

main();


