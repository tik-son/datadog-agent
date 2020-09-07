require 'spec_helper'

describe 'successfully run functional test' do
  it 'displays PASS and returns 0' do
    output = `sudo /tmp/security-agent/testsuite`
    expect($?).to eq(0)
    expect(output).to include("PASS")
    print output
  end
end
