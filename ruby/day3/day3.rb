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
    readings = read_input_file_lines()
    # readings = ["00100",
    # "11110",
    # "10110",
    # "10111",
    # "10101",
    # "01111",
    # "00111",
    # "11100",
    # "10000",
    # "11001",
    # "00010",
    # "01010"]

    reduced = work_on(readings, 0)

    reduced2 = work_on2(readings, 0)
    return bin_to_dec(reduced) * bin_to_dec(reduced2)
end

def work_on(readings, bit_index) 
    # puts "readings: #{readings}"

    if readings.length == 1
        return readings[0]
    end

    most_common = get_most_common_bits(readings)[bit_index]
    # puts "most_common: #{most_common}"
    filtered_readings = readings.filter { |r| r[bit_index] == most_common }
    # puts "filtered_readings: #{filtered_readings}"

    return work_on(filtered_readings, bit_index + 1)
end

def work_on2(readings, bit_index)
    # puts "readings: #{readings}"

    if readings.length == 1
        return readings[0]
    end

    least_common = get_least_common_bits(readings)[bit_index]
    # puts "least_common: #{least_common}"
    filtered_readings = readings.filter { |r| r[bit_index] == least_common }
    # puts "filtered_readings: #{filtered_readings}"
    # puts "========="
    return work_on2(filtered_readings, bit_index + 1)
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

def get_least_common_bits(readings)
    width = readings[0].length
    positional_bits = []
    (0..width-1).each do |i|
        bap = get_bits_at_position(readings, i)
        positional_bits.push(bap)
    end
    
    least_common = find_least_common(positional_bits)
    return least_common
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
        if p.count("1") * 2  >= p.length
            most_common += "1"
        else
            most_common += "0"
        end
    end
    return most_common
end

def find_least_common(positional_bits) 
    least_common = ""
    positional_bits.each do |p| 
        if p.count("0") * 2  <= p.length
            least_common += "0"
        else
            least_common += "1"
        end
    end
    return least_common
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