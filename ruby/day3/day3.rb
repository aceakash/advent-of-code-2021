def part_one()
    readings = read_input_file_lines()
    # readings = ["10", "11", "01"]

    most_common_bits = get_most_common_bits(readings) # 10111
    least_common_bits = invert(most_common_bits)

    gamma = bin_to_dec(most_common_bits)
    epsilon = bin_to_dec(least_common_bits)

    return gamma * epsilon
end

def part_two()
    # readings = read_input_file_lines()
    readings = ["10", "11", "01"]

    pos = 0
    r1 = readings.dup
    while r1.length > 1
        bits = get_bits_at_position(r1, pos)
        most_common_bit = find_most_common(bits)
        r1 = only_keep_starting_with(r1, most_common_bit)
        pos += 1
    end

    return r1[0]
end

def only_keep_starting_with(readings, bit_char)
    return readings.select {|r| r.start_with? bit_char } 
end

def get_most_common_bits(readings)
    width = readings[0].length
    positional_bits = []
    (0..width-1).each do |i|
        bap = get_bits_at_position(readings, i)
        positional_bits.push(bap)
    end
    
    most_common = find_most_common(positional_bits)
    return most_common
end

def get_bits_at_position(readings, pos) 
    bits = ""
    readings.each do |r|
        bits += r[pos]
    end
    return bits
end

def find_most_common(positional_bits) 
    most_common = ""
    positional_bits.each do |p| 
        if p.count("1") > p.length/2
            most_common += "1"
        else
            most_common += "0"
        end
    end
    return most_common
end

def invert(bits)
    inverted = bits.chars.map do |c| 
        c == "0" ? "1" : "0"
    end
    return inverted.join ""
end

def bin_to_dec(str) 
    return str.to_i(2)
end

def read_input_file_lines()
    file = File.open("input.txt")
    file_data = file.readlines.map(&:chomp)
    return file_data
end


puts "Part one: #{part_one()}"

puts "Part two: #{part_two()}"