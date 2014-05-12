h = Hash.new 0
while gets
  $_.chomp.split(/\s+/).each{|w| 
    if w =~ /^[\u0E00-\u0EFF]+$/
      h[w] += 1
    end
  }
end

h.each{|k,v| puts "#{k}\t#{v}"}
