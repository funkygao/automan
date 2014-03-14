<?php

class JsonTag extends Annotation
{

    /**
     * Json content array.
     *
     * If starts with '@', include external json file.
     * Else, content is json string itself.
     */
    public function getJsonArray($filename = '')
    {
        if ($this->value[0] == '@') {
            $jsonFile = dirname($filename) . '/'
                . substr($this->value, -(strlen($this->value)-1));
            $this->value = file_get_contents($jsonFile);
        }

        return json_decode($this->value, TRUE);
    }

}
