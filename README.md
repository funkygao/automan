iodoc
=====

extension for phpDocumentor that can support @input @output tag

Features
========

*   helps frontend/backend/pm all
    - parallel dev
    - contract by auto generated doc
*   auto generate api documentation
*   auto generate runnable backend integration test scripts
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
           |----+                                   |           |
           |    |                                   |           |
           |    | write php skeleton                |           |
           |    | @input/@output tags               |           |
           |    |                                   |           |
           |<---+                                   |           |
           |                                        |           |
           |                                        |           |
           |  3. export generated api doc           |           |
           |--------------------------------------->|---------->|
           |                                        |           |
           |                                        |           |
           |----+                                   |           |
           |    |                                   |           |
           |    | run generated test scripts        |           |
           |    |                                   |           |
           |<---+                                   |           |
           |                                        |           |
           |                                        |           |


