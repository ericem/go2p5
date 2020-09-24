# -*- mode: ruby -*-
# vi: set ft=ruby :

VAGRANT_ROOT = File.dirname(File.expand_path(__FILE__))
disk_file = File.join(VAGRANT_ROOT, 'disk.vdi')

Vagrant.configure("2") do |config|
  config.vm.box = "oraclelinux/7"

  config.vm.box_url = "https://oracle.github.io/vagrant-projects/boxes/oraclelinux/7.json"

  config.vm.provider "virtualbox" do |vb|
    vb.name = "twofive"
    vb.memory = 1024
    vb.cpus = 2
    unless File.exist?(disk_file)
      vb.customize [ "createhd", "--filename", disk_file, "--size", "102400"]
    end
    vb.customize [ "storageattach", "twofive" , "--storagectl",
                   "SATA Controller", "--port", "1", "--device", "0",
                   "--type",  "hdd", "--medium", disk_file ]
  end
end
