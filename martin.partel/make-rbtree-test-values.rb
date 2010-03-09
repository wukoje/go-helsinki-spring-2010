#!/usr/bin/env ruby

amount = 100
max_key = 1000

test_values = (0...amount).map { rand(max_key) }

outside_values = []
while outside_values.size < amount / 2
	k = rand(max_key)
	outside_values << k unless test_values.include? k
end

values_in_delete_order = test_values.shuffle

duplicates = []
test_values.each do |x|
	duplicates << x if test_values.count(x) > 1
end

puts "\ttestValues := []int{#{test_values.join ', '}}"
puts "\toutsideValues := []int{#{outside_values.join ', '}}"
puts "\tvaluesInDeleteOrder := []int{#{values_in_delete_order.join ', '}}"
puts "\t// duplicates: []int{#{duplicates.join ', '}}"
