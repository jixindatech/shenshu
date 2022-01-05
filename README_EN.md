# shenshu
English | [中文](./README.md)
web firewall admin

## install
1. Web component adapts vue，server use golang，install easily.

## config introduction
1. auth.json based on rbac. only admin role exists，others could be added like admin easily.
2. basic_model.conf and basic_policy.csv based on casbin for api authority check, others could be added like admin easily.
3. about config.yaml，all configuration stored on mysql，redis data stored for shenshu gateway and
shenshu admin,just for data exchanged，all rule events ard stored on elasticsearch.
configuration about redis and elasticsearch should be the same with shenshu gateway.

##Contributing
you are wellcome for issue and star

## Discussion Group
QQ group: 254210748

## License
Unlicense



