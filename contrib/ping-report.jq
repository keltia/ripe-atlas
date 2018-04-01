map(.result) | flatten(1) | map(.rtt) | length as $total |
 "Median: " + (sort |
      if length % 2 == 0 then .[length/2] else .[(length-1)/2] end | tostring),
 "Average: " + (map(select(. != null)) | add/length | tostring) + " ms",
 "Min: " + (map(select(. != null)) | min | tostring) + " ms",
 "Max: " + (max | tostring) + " ms",
 "Failures: " + (map(select(. == null)) | (length*100/$total) | tostring) + " %"
