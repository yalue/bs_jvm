# This file is used to convert an HTML table into go function stubs and tables.

# Inserts a newline character into s so that the first line will not exceed 80
# characters. The newline will be placed after the last possible comma without
# exceeding the 80 character limit. (Intended to be used for long function
# definitions).
def split_line(s)
  return s if s.size <= 80
  parts = s.split(",")
  to_return = ""
  while parts.size > 0
    if (to_return.size + parts[0].size + 1) >= 80
      to_return += ",\n"
      break
    end
    if to_return.size == 0
      to_return += parts.shift
    else
      to_return += "," + parts.shift
    end
  end
  to_return += "\t" + parts.join(",").sub(/^\s+/, "")
  to_return
end

class JVMOpcode
  attr_accessor :opcode, :name, :has_extra_bytes

  def to_s
    has_extra = ""
    has_extra = " + extra bytes" if @has_extra_bytes
    "%s: 0x%02x%s" % [@name, @opcode, has_extra]
  end

  # Returns the name of the parser function in go.
  def go_parser_function_name
    "parse#{@name.capitalize}Instruction"
  end

  # Returns the opcode table
  def to_go_struct
    parser_name = self.go_parser_function_name
    to_return = "&jvmOpcodeInfo{\n"
    to_return += "\tname: \"#{@name}\",\n"
    to_return += "\topcode: 0x%02x,\n" % [@opcode]
    to_return += "\tparse: #{parser_name},\n"
    to_return += "}"
  end

  def to_go_type
    "#{@name}Instruction"
  end

  def to_go_type_definition
    "type #{self.to_go_type} struct{ knownJVMInstruction }\n"
  end

  def to_go_parser_function
    signature = "func #{self.go_parser_function_name}("
    signature += "opcode uint8, name string, address uint, m JVMMemory) "
    signature += "(JVMInstruction, error) {"
    signature = split_line(signature)
    to_return = signature + "\n"
    to_return += "\ttoReturn := #{self.to_go_type}{\n"
    to_return += "\t\tknownJVMInstruction{\n"
    to_return += "\t\t\traw: 0x%02x,\n" % [@opcode]
    to_return += "\t\t\tname: name,\n"
    to_return += "\t\t},\n"
    to_return += "\t}\n"
    to_return += "\treturn &toReturn, nil\n"
    to_return += "}\n"
    to_return
  end

  # Returns a parser function stub for the given instruction, which returns a
  # NotImplementedError
  def generate_parser_stub
    signature = "func #{self.go_parser_function_name}("
    signature += "opcode uint8, name string, address uint, m JVMMemory) "
    signature += "(JVMInstruction, error) {"
    signature = split_line(signature)
    to_return = signature + "\n"
    to_return += "\treturn nil, NotImplementedError\n"
    to_return += "}\n"
    to_return
  end

  def generate_execute_stub
    signature = "func (n *#{self.to_go_type}) Execute(t JVMThread) error {"
    signature = split_line(signature)
    to_return = signature + "\n"
    to_return += "\treturn nil, NotImplementedError\n"
    to_return += "}\n"
    to_return
  end
end

# Returns a go map of the instructions
def generate_go_opcode_table(instructions)
  to_return = "var opcodeTable = map[uint8]*jvmOpcodeInfo{\n"
  instructions.each {|n| to_return += n.to_go_struct + ",\n"}
  to_return += "}"
  to_return
end

# Generates go code for type definitions and parser functions. Doesn't generate
# types for multi-byte opcodes.
def generate_go_parsers_and_types(instructions)
  to_return = ""
  instructions.each do |n|
    if n.has_extra_bytes
      to_return += n.generate_parser_stub + "\n"
      next
    end
    to_return += n.to_go_type_definition + "\n"
    to_return += n.to_go_parser_function + "\n"
  end
  to_return
end

# Generates stubs for Execute() functions for each instruction.
def generate_go_execute_stubs(instructions)
  to_return = ""
  instructions.each do |n|
    to_return += n.generate_execute_stub + "\n"
  end
  to_return
end

content = ""
File.open("java_bytecode.txt", 'rb') {|f| content = f.read.gsub(/\r?\n/, " ")}
content = content.scan(/<tr>.*?<\/tr>/)
content = content.map do |row|
  cells = row.scan(/<td>.*?<\/td>/)
  cells = cells.map do |cell|
    content = ""
    if cell =~ /<td>(.*?)<\/td>/
      content = $1
    end
    content
  end
  cells[1] = cells[1].to_i(16)
  instruction = JVMOpcode.new
  instruction.opcode = cells[1]
  instruction.name = cells[0]
  instruction.has_extra_bytes = !!(cells[3] =~ /\S/)
  instruction
end
content.sort! {|a, b| a.opcode <=> b.opcode}

puts " Opcode table ".center(80, "#")
puts generate_go_opcode_table(content)
puts " Parser functions and types ".center(80, "#")
puts generate_go_parsers_and_types(content)
puts " Execute function stubs ".center(80, "#")
puts generate_go_execute_stubs(content)
