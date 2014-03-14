automan
=======

Auto generate api doc for backend/frontend developpers for collaboration.
    
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
*   helps frontend/backend/PM all
    - parallel development
    - contract by auto generated doc
*   keep doc and src close enough
    - less maintainance cost
    - better readability


Usage
=====

        BackendDeveloper                FrontendDeveloper       PM
        ----------------                -----------------       --
           |                                        |           |
           |  1. design the call interface          |           |
           |<-------------------------------------->|           |
           |                                        |           |
           |-----+                                  |           |
           |     |                                  |           |
           |  2. | write php skeleton               |           |
           |     | @input/@output tags              |           |
           |     |                                  |           |
           |<----+                                  |           |
           |                                        |           |
           |                                        |           |
           |  3. export generated api doc           |           |
           |--------------------------------------->|---------->|
           |                                        |           |
           |                                        |           |
           |-----+                                  |           |
           |     |                                  |           |
           |  4. | run generated test scripts       |           |
           |     |                                  |           |
           |<----+                                  |           |
           |                                        |           |
           |                                        |           |


