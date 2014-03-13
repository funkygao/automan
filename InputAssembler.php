<?php
/**
 * phpDocumentor
 *
 * PHP Version 5.3
 *
 * @copyright 2010-2014 Mike van Riel / Naenius (http://www.naenius.com)
 * @license   http://www.opensource.org/licenses/mit-license.php MIT
 * @link      http://phpdoc.org
 */

namespace phpDocumentor\Descriptor\Builder\Reflector\Tags;

use phpDocumentor\Descriptor\Builder\Reflector\AssemblerAbstract;
use phpDocumentor\Descriptor\Tag\ParamDescriptor;
use phpDocumentor\Reflection\DocBlock\Tag\InputTag;

class InputAssembler extends AssemblerAbstract
{
    /**
     * Creates a new Descriptor from the given Reflector.
     *
     * @param InputTag $data
     *
     * @return InputDescriptor
     */
    public function create($data)
    {
        $descriptor = new InputDescriptor($data->getName());
        $descriptor->setDescription($data->getDescription());
        //$descriptor->setVariableName($data->getVariableName());
        $descriptor->setTypes($data->getTypes());

        return $descriptor;
    }
}
