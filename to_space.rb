while gets
  puts $_.chomp.gsub(/:|`|<[^>]+>|{{|}}|\[\[|\]\]|\||\/|'|=/, ' ')
end
