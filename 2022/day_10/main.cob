IDENTIFICATION DIVISION.
PROGRAM-ID. DAY_10.
AUTHOR. mrwormhole.
DATE-WRITTEN. 10th of December, 2022.

ENVIRONMENT DIVISION.
   INPUT-OUTPUT SECTION.
      FILE-CONTROL.
      SELECT INPUT-FILE ASSIGN TO 'input.txt'
      ORGANIZATION IS LINE SEQUENTIAL.          

DATA DIVISION.
   FILE SECTION.
   FD INPUT-FILE.
   01 INPUT-RECORD.
      05 INPUT-LINE  PIC X(100).

   WORKING-STORAGE SECTION.
   01 WS-INPUT-RECORD.
      05 WS-INPUT-LINE  PIC X(100).
   01 WS-EOF PIC A(1). 
   01 WS-MINUS-COUNT PIC 9(1).
   01 WS-COMMAND PIC X(4).
   01 WS-VALUE PIC S9(32) VALUE ZEROES.
   01 WS-CYCLE PIC 9(32) VALUE 1.
   01 WS-TOTAL PIC S9(32) VALUE 1.
   01 WS-TEMP PIC S9(32) VALUE ZEROES.
   01 WS-SIGNAL-STRENGTH PIC S9(32) VALUE ZEROES.
   01 WS-OVER20 PIC A(1) VALUE 'F'.
   01 WS-OVER60 PIC A(1) VALUE 'F'.
   01 WS-OVER100 PIC A(1) VALUE 'F'.
   01 WS-OVER140 PIC A(1) VALUE 'F'.
   01 WS-OVER180 PIC A(1) VALUE 'F'.
   01 WS-OVER220 PIC A(1) VALUE 'F'.

PROCEDURE DIVISION.
   OPEN INPUT INPUT-FILE.
      PERFORM UNTIL WS-EOF='Y'
         READ INPUT-FILE 
            INTO WS-INPUT-RECORD
            AT END 
               MOVE 'Y' TO WS-EOF
            NOT AT END 
               EVALUATE WS-INPUT-LINE
                   WHEN "noop"
                      COMPUTE WS-CYCLE = WS-CYCLE + 1
                   WHEN OTHER
                      MOVE 0 TO WS-MINUS-COUNT
                      INSPECT WS-INPUT-LINE 
                           TALLYING WS-MINUS-COUNT 
                           FOR ALL "-"
                      IF WS-MINUS-COUNT > 0 THEN 
                        UNSTRING WS-INPUT-LINE 
                           DELIMITED BY "-"
                           INTO WS-COMMAND
                                WS-VALUE
                        END-UNSTRING    
                        COMPUTE WS-VALUE = WS-VALUE * -1
                        COMPUTE WS-TOTAL = WS-TOTAL + WS-VALUE 
                      ELSE 
                        UNSTRING WS-INPUT-LINE 
                           DELIMITED BY " "
                           INTO WS-COMMAND
                                WS-VALUE
                        END-UNSTRING
                        COMPUTE WS-TOTAL = WS-TOTAL + WS-VALUE
                      END-IF
                      COMPUTE WS-CYCLE = WS-CYCLE + 2
               END-EVALUATE
               IF WS-CYCLE - 20 < 2 AND WS-CYCLE - 20 > -1 AND WS-OVER20 = 'F' THEN  
                   IF WS-CYCLE = 20 THEN 
                       COMPUTE WS-TEMP = 20 * WS-TOTAL
                   ELSE
                       COMPUTE WS-TEMP = 20 * (WS-TOTAL - WS-VALUE)
                   END-IF
                   COMPUTE WS-SIGNAL-STRENGTH = WS-SIGNAL-STRENGTH + WS-TEMP
                   MOVE 'T' to WS-OVER20
               END-IF
               IF WS-CYCLE - 60 < 2 AND WS-CYCLE - 60 > -1 AND WS-OVER60 = 'F' THEN  
                   IF WS-CYCLE = 60 THEN 
                       COMPUTE WS-TEMP = 60 * WS-TOTAL
                   ELSE
                       COMPUTE WS-TEMP = 60 * (WS-TOTAL - WS-VALUE)
                   END-IF
                   COMPUTE WS-SIGNAL-STRENGTH = WS-SIGNAL-STRENGTH + WS-TEMP
                   MOVE 'T' to WS-OVER60
               END-IF
               IF WS-CYCLE - 100 < 2 AND WS-CYCLE - 100 > -1 AND WS-OVER100 = 'F' THEN  
                   IF WS-CYCLE = 100 THEN 
                       COMPUTE WS-TEMP = 100 * WS-TOTAL
                   ELSE
                       COMPUTE WS-TEMP = 100 * (WS-TOTAL - WS-VALUE)
                   END-IF
                   COMPUTE WS-SIGNAL-STRENGTH = WS-SIGNAL-STRENGTH + WS-TEMP
                   MOVE 'T' to WS-OVER100
               END-IF
               IF WS-CYCLE - 140 < 2 AND WS-CYCLE - 140 > -1 AND WS-OVER140 = 'F' THEN  
                   IF WS-CYCLE = 140 THEN 
                       COMPUTE WS-TEMP = 140 * WS-TOTAL
                   ELSE
                       COMPUTE WS-TEMP = 140 * (WS-TOTAL - WS-VALUE)
                   END-IF
                   COMPUTE WS-SIGNAL-STRENGTH = WS-SIGNAL-STRENGTH + WS-TEMP
                   MOVE 'T' to WS-OVER140
               END-IF
               IF WS-CYCLE - 180 < 2 AND WS-CYCLE - 180 > -1 AND WS-OVER180 = 'F' THEN  
                   IF WS-CYCLE = 180 THEN 
                       COMPUTE WS-TEMP = 180 * WS-TOTAL
                   ELSE
                       COMPUTE WS-TEMP = 180 * (WS-TOTAL - WS-VALUE)
                   END-IF
                   COMPUTE WS-SIGNAL-STRENGTH = WS-SIGNAL-STRENGTH + WS-TEMP
                   MOVE 'T' to WS-OVER180
               END-IF
               IF WS-CYCLE - 220 < 2 AND WS-CYCLE - 220 > -1 AND WS-OVER220 = 'F' THEN
                   IF WS-CYCLE = 220 THEN 
                       COMPUTE WS-TEMP = 220 * WS-TOTAL
                   ELSE
                       COMPUTE WS-TEMP = 220 * (WS-TOTAL - WS-VALUE)
                   END-IF
                   COMPUTE WS-SIGNAL-STRENGTH = WS-SIGNAL-STRENGTH + WS-TEMP
                   MOVE 'T' to WS-OVER220
               END-IF
         END-READ
      END-PERFORM.
      DISPLAY "PART1: " WS-SIGNAL-STRENGTH.
   CLOSE INPUT-FILE.
STOP RUN.


