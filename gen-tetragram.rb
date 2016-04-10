#!/usr/bin/env ruby

monograms = {0 => "------",
             1 => "--  --",
             2 => "- - - ",
             }
#(1..4).reduce([]) {|acc, i| acc << rand(3).to_i; acc}.map {|i| puts "#{monograms[i]}\n"}
tetragram = (1..4).reduce([]) do |acc, i|
  acc << rand(3).to_i
  acc
end

# print out the lines
tetragram.map {|i| puts "#{monograms[i]}\n"}

# show the Passage number
base10_sum = 0
tetragram.each_with_index do |value, index|
  case index
     when 0 then base10_sum += value*27
     when 1 then base10_sum += value*9
     when 2 then base10_sum += value*3
     when 3 then base10_sum += value*1
  end
end
puts "\nPassage: #{base10_sum + 1}"
