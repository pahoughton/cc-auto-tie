# 2018-10-13 (cc) <paul4hough@gmail.com>
#

# 2018-10-10 (cc) <paul4hough@gmail.com>
#
# y = rake recurses down (. .. ../..:)

$runstart = Time.now

at_exit {
  runtime = Time.at(Time.now - $runstart).utc.strftime("%H:%M:%S.%3N")
  puts "run time: #{runtime}"
}

task :default do
  sh 'rake --tasks'
  exit 1
end

task :syntax do
  sh 'ansible-playbook --syntax-check ansible/maul.yml'
end


task :lint do
  sh "yamllint -f parsable ."
end


task :provision do |task, args|
  sh "vagrant provision"
end

task :rebuild do |task, args|
  sh "vagrant destroy -f"
  sh "vagrant up"
end
