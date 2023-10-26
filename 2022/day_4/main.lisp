(defvar *filename* "input.txt")
(defvar first_count 0)
(defvar second_count 0)

(defun delimiterc (c) (char= c #\,))

(defun split_via_comma (string &key (delimiterc #'delimiterc))
  (loop :for beg = (position-if-not delimiterc string)
    :then (position-if-not delimiterc string :start (1+ end))
    :for end = (and beg (position-if delimiterc string :start beg))
    :when beg :collect (subseq string beg end)
    :while end
  )
)

(defun delimeterh (c) (char= c #\-))

(defun split_via_hyphen (string &key (delimeterh #'delimeterh))
  (loop :for beg = (position-if-not delimeterh string)
    :then (position-if-not delimeterh string :start (1+ end))
    :for end = (and beg (position-if delimeterh string :start beg))
    :when beg :collect (subseq string beg end)
    :while end
  )
)

(defun range (min max &optional (step 1))
  (when (<= min max)
    (cons min (range (+ min step) max step))
  )
)

(defun is_contained (a1 a2 b1 b2)
    (setq once nil)

    (if (and (>= a1 b1) (<= a2 b2))
      (progn
        (setq first_count (+ first_count 1))
        (setq once T)
      )
    )

    (if (and (>= b1 a1) (<= b2 a2))
      (progn
        (if (eq once nil)
            (setq first_count (+ first_count 1))
        )
      )
    )
)

(defun is_overlapped (a1 a2 b1 b2)
    (loop named outer for a in (range a1 a2) do
      (loop for b in (range b1 b2) do
        (if (eq a b)
          (progn
            (setq second_count (+ second_count 1))
              (return-from outer)
            )
        )
      )
    )
)

(let ((in (open *filename* :if-does-not-exist nil)))
   (when in
      (loop for line = (read-line in nil) while line do
        (format t "~a~%" line)

        (setq first_interval (nth 0 (split_via_comma line)))
        (setq second_interval (nth 1 (split_via_comma line)))
        (setq first_interval_start (parse-integer (nth 0 (split_via_hyphen first_interval))))
        (setq first_interval_end (parse-integer (nth 1 (split_via_hyphen first_interval))))
        (setq second_interval_start (parse-integer (nth 0 (split_via_hyphen second_interval))))
        (setq second_interval_end (parse-integer (nth 1 (split_via_hyphen second_interval))))
    
        (is_contained first_interval_start first_interval_end second_interval_start second_interval_end)
        (is_overlapped first_interval_start first_interval_end second_interval_start second_interval_end)
      )
      (close in)
   )
)

(format t "first_count ~a~%" first_count)
(format t "second_count ~a~%" second_count)