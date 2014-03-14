<?php

class JsonTag extends Annotation
{

    /**
     * Json content.
     *
     * If starts with '@', include external json file.
     * Else, content is json string itself.
     *
     * @var string
     */
    public $json = NULL;

    public function getJsonArray($filename = '')
    {
        if ($this->json[0] == '@') {
            $jsonFile = dirname($filename) . '/'
                . substr($this->json, -(strlen($this->json)-1));
            $this->json = file_get_contents($jsonFile);
        }

        return json_decode($this->json, TRUE);
    }

}