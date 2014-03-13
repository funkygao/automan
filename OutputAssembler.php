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
use phpDocumentor\Descriptor\Tag\OutputDescriptor;
use phpDocumentor\Reflection\DocBlock\Tag\OutputTag;

class OutputAssembler extends AssemblerAbstract
{
    /**
     * Creates a new Descriptor from the given Reflector.
     *
     * @param OutputTag $data
     *
     * @return OutDescriptor
     */
    public function create($data)
    {
        $descriptor = new OutDescriptor($data->getName());
        $descriptor->setDescription($data->getDescription());
        //$descriptor->setVariableName($data->getVariableName());
        $descriptor->setTypes($data->getTypes());

        return $descriptor;
    }
}
