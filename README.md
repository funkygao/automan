automan
=======

Auto generate [api doc | mock service/data | test script] for backend/frontend developpers for loosely coupled teamwork.
    
                   __                                
    _____   __ ___/  |_  ____   _____ _____    ____  
    \__  \ |  |  \   __\/  _ \ /     \\__  \  /    \ 
     / __ \|  |  /|  | (  <_> )  Y Y  \/ __ \|   |  \
    (____  /____/ |__|  \____/|__|_|  (____  /___|  /
         \/                         \/     \/     \/ 
    

Features
========

*   auto generate api documentation
*   auto generate runnable backend integration test scripts
*   auto generate web server that can reply to frontend request according to tags definition
    - generate runnable mock service
    - generate mock data
*   helps frontend/backend/PM all
    - parallel development
    - contract by auto generated doc
*   keep doc and src close enough
    - less maintainance cost
    - better readability


Workflow
========

        BackendDeveloper                FrontendDeveloper       PM
        ----------------                -----------------       --
           |                                        |           |
           |  1. design the call interface          |           |
           |<-------------------------------------->|<--------->|
           |                                        |           |
           |-----+                                  |           |
           |     |                                  |           |
           |  2. | write php skeleton               |           |
           |     | with @In/@Out tags               |           |
           |     |                                  |           |
           |<----+                                  |           |
           |                                        |           |
           |                                        |           |
           |  3. export generated api doc           |           |
           |--------------------------------------->|---------->|
           |     wiki                               |           |
           |                                        |      
           |-----+                                  |-----+
           |     |                                  |     |
           |  4. | run generated test scripts       |  4. | connect flash/Xcode to      
           |     |                                  |     | mock server
           |<----+                                  |<----+   
           |                                        |        
           |                                        |       


Usage
=====


### 1. php integration

    class UserService {
    
        /**
         * Input/Output is referenced by external file.
         * Cause its values starts with '@'.
         *
         * @Spec('http://wiki.mycorp.com/display/DG/test+auto+wiki')
         * @In('@json/hello.i.json')
         * @Out('@json/hello.o.json')
         */
        public function helloWorld($param) {
            return array(
                'uid' => 5,
                'status' => 'ok',
            );
        }
    
        /**
         * Raw json string.
         *
         * @In('{"uid":3}')
         * @Out('{"status":"ok","time":139343434,"msg":"good"}')
         */
        public function foo($param) {
        }
    }

### 2. generate stub/mock

    MODIFY bootstrap.php before you use automan!

    nohup ./webmocker/webmocker&               # this is the dumb web server for mock services
    ./automan ${your_php_controller_file_path} # auto kill -HUP webmocker; auto generate test.sh

    e.g, ./automan /Users/gaopeng/fun/dragon-server-code/v2/classes/Services/AccountService.php

### 3. integration with conflucnece rest api

    ./autowiki output/automan.api

