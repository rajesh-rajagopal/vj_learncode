require 'nokogiri'

@doc = Nokogiri::XML(File.open("show.xml"))
@context = @doc.xpath("//CONTEXT")
puts @context.xpath("//ACCOUNT_ID")
puts "-------------------"
puts @doc.xpath("//ACCOUNTS_ID")
puts @doc.xpath("//ASSEMBLY_ID")
puts @doc.xpath("//ASSEMBLIES_ID")

@doc.css("CONTEXT").each do |response_node|
  puts response_node.text if response_node.name == "ASSEMBLY_ID"
end
