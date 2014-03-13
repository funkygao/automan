iodoc
=====

extension for phpDocumentor that can support @input @output tag

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


