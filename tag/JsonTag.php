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
        if ($this->value && $this->value[0] == '@') {
            $jsonFile = dirname($filename) . '/'
                . substr($this->value, -(strlen($this->value)-1));
            $this->value = file_get_contents($jsonFile);
        }

        $arr = json_decode($this->value, TRUE);
        if (json_last_error() != JSON_ERROR_NONE) {
            echo "Invalid tag defined: \n";
            print_r($this);
        }
        if (is_array($arr)) {
            return $arr;
        } else {
            return $this->value;
        }
    }

}
