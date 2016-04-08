#!/usr/bin/env ruby

monograms = {0 => "------",
             1 => "--  --",
             2 => "- - - ",
             }
(1..4).reduce([]) {|acc, i| acc << rand(3).to_i; acc}.map {|i| puts "#{monograms[i]}\n"}
