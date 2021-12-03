def part_one()
    readings = read_input_file_lines()
    # readings = ["10", "11", "01"]

    most_common_bits = get_most_or_least_common_bits(readings, :find_most_common) # 10111
    least_common_bits = invert(most_common_bits)

    gamma = bin_to_dec(most_common_bits)
    epsilon = bin_to_dec(least_common_bits)

    return gamma * epsilon
end

def part_two()
    readings = read_input_file_lines()
    most_common = process(readings, 0, "most")
    least_common = process(readings, 0, "least")

    return bin_to_dec(most_common) * bin_to_dec(least_common)
end

def process(readings, bit_index, most_or_least) 
    if readings.length == 1
        return readings[0]
    end

    func = most_or_least == "most" ? :find_most_common : :find_least_common

    most_common = get_most_or_least_common_bits(readings, func)[bit_index]
    filtered_readings = readings.filter { |r| r[bit_index] == most_common }
    return process(filtered_readings, bit_index + 1, most_or_least)
end

def only_keep_starting_with(readings, bit_char)
    return readings.select {|r| r.start_with? bit_char } 
end

def get_most_or_least_common_bits(readings, find_least_or_most_common)
    width = readings[0].length
    positional_bits = []
    (0..width-1).each do |i|
        bap = get_bits_at_position(readings, i)
        positional_bits.push(bap)
    end
    
    least_or_most_common = method(find_least_or_most_common).call(positional_bits)
    return least_or_most_common
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