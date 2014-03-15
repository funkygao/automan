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
           |                                        |           |
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

Modify bootstrap.php before you use automan.

    class UserService {
    
        /**
         * Input/Output is referenced by external file.
         * Cause its values starts with '@'.
         *
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
         */
        public function foo($param) {
        }
    }

