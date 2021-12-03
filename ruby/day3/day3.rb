def part_one()
    readings = read_input_file_lines()[0]

    most_common_bits = get_most_common_bits(readings) # 10111
    least_common_bits = invert(most_common_bits)

    gamma = bin_to_dec(most_common_bits)
    epsilon = bin_to_dec(least_common_bits)

    return gamma * epsilon
end

def read_input_file_lines()
    file = File.open("input.txt")
    file_data = file.readlines.map(&:chomp)
    return file_data
end


puts "Part one: #{part_one()}"