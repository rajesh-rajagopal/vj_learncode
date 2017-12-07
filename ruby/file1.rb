File.open("test-file", "w") do |file|
  file.puts "Hello file!"
end

####################################

f = File.open('some-file.txt', 'a')

f.write('here is some text')


f.close


output = File.open( "outputfile.yml","w" )
output << "This is going to the output file"
output.close
