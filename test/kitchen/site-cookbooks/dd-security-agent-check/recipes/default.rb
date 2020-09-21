#
# Cookbook Name:: dd-security-agent-check
# Recipe:: default
#
# Copyright (C) 2020 Datadog
#

if node['platform_family'] != 'windows'
  wrk_dir = '/tmp/security-agent'

  package 'Install i386 libc' do
    case node[:platform]
    when 'redhat', 'centos', 'suse', 'fedora'
      package_name 'libc.i686'
    when 'ubuntu', 'debian'
      package_name 'libc6-i386'
    end
  end

  kernel_module 'loop' do
    action :load
  end

  directory wrk_dir do
    recursive true
  end

  cookbook_file "#{wrk_dir}/testsuite" do
    source "testsuite"
    mode '755'
  end

  cookbook_file "#{wrk_dir}/testsuite32" do
    source "testsuite32"
    mode '755'
  end
end