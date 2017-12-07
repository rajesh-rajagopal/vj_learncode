require 'nokogiri'

@doc = Nokogiri::XML(File.open("show-nfs.xml"))

DISK_TYPE = @doc.xpath("//DISK_TYPE").text

if DISK_TYPE == "FILE" 
  @doc.xpath('//MONITORING//DISK_SIZE/*').each do |row|
     puts row.xpath('//ID')
  end
end

#puts "SIZE :#{SIZE}"


