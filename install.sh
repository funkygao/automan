#!/bin/sh
#==================================
# install this for phpDocumentor
#==================================

BASE=/opt/app/pear/share/pear/phpDocumentor/src/phpDocumentor

install InputAssembler.php $BASE/Descriptor/Builder/Reflector/Tags/InputAssembler.php
install InputTag.php $BASE/Plugin/Core/Transformer/Behaviour/Tag/InputTag.php
install OutputAssembler.php $BASE/Descriptor/Builder/Reflector/Tags/OutputAssembler.php
install OutputTag.php $BASE/Plugin/Core/Transformer/Behaviour/Tag/OutputTag.php

echo 'Please patch phpDocumentor manually!!!'
