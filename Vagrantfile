# 2018-12-23 (cc) <paul4hough@gmail.com>
#

Vagrant.configure("2") do |config|
  # config.vm.box_check_update = false

  pgsql = 'pgsql'
  config.vm.define pgsql do |bcfg|
    bcfg.vm.box = "centos/7"
    # bcfg.vm.box = "ubuntu/xenial64"

    bcfg.vm.hostname = pgsql
    bcfg.vm.network    'private_network', ip: '10.0.0.6'
    bcfg.vm.provider   'virtualbox' do |vb|
      vb.name      = pgsql
      vb.cpus      = 1
      vb.memory    = 1024
      vb.customize ['modifyvm', :id, '--natdnshostresolver1', 'on']
      vb.customize ['modifyvm', :id, '--natdnspassdomain1', 'on']
      vb.customize ['modifyvm', :id, '--usb', 'off']
    end
    bcfg.vm.provision "ansible" do |ansible|
      ansible.playbook = "ansible/node-pgsql.yml"
    end
  end

  prom = 'prom'
  config.vm.define prom do |bcfg|
    bcfg.vm.box = "centos/7"
    # bcfg.vm.box = "ubuntu/xenial64"

    bcfg.vm.hostname = prom
    bcfg.vm.network    'private_network', ip: '10.0.0.5'
    bcfg.vm.provider   'virtualbox' do |vb|
      vb.name      = prom
      vb.cpus      = 1
      vb.memory    = 1024
      vb.customize ['modifyvm', :id, '--natdnshostresolver1', 'on']
      vb.customize ['modifyvm', :id, '--natdnspassdomain1', 'on']
      vb.customize ['modifyvm', :id, '--usb', 'off']
    end
    bcfg.vm.provision "ansible" do |ansible|
      ansible.playbook = "ansible/node-prom.yml"
    end
  end

end
