import React, { Component } from 'react';
import './index.css';
import { Input, Radio, Table, Divider, Tag, message, Button } from 'antd';
import request from '../../utils/request';
import ConfigModal from './modal/configModal';

const Search = Input.Search;
class ConfigList extends Component {
    constructor() {
        super()
        this.state = {
            loading: true,
            modalType: 'create',
            modalVisible: false,
            modalData: {},
        }
    }
    componentDidMount() {
        this.loadData()
    }
    async loadData(){
        this.setState({
            loading:true
        })
        const res = await request.get("/config")
        if (res.status != 200) {
            message.error("get config error")
            return
        }
        const configObj = res.data;
        let configArr = [];
        for (const key in configObj) {
            if (configObj.hasOwnProperty(key)) {
                configArr.push({
                    key: key,
                    value: configObj[key],
                })
            }
        }
        this.setState({
            loading: false,
            data: configArr
        })
    }
    handleCancel() {
        this.setState({
            modalVisible: false
        })
    }
    handleCreateConfig() {
        this.setState({
            modalVisible: true,
            modalData: {}
        })
    }

    handleEditConfig(data) {
        this.setState({
            modalVisible: true,
            modalData: data
        })
    }
    async handleDeleteConfig(key) {
        const res = await request.post("/config/delete", { key: key })
        if (res.status != 200) {
            message.error("删除失败")
            return
        }
        message.success("删除成功")
        this.loadData()
    }

    render() {
        const columns = [{
            title: 'Key',
            dataIndex: 'key',
            key: 'key',
            width: 100,
            render: text => <a href="javascript:;">{text}</a>,
        }, {
            title: 'Value',
            key: 'value',
            dataIndex: 'value',
            width: 200,
            render: value => (
                <span>
                    {value}
                </span>
            ),
        }, {
            title: 'Action',
            key: 'action',
            width: 50,
            render: (text, record) => (
                <div>
                    <span onClick={() => this.handleEditConfig(record)}>
                        <a href="javascript:;">详情</a>
                    </span>
                    <span>  </span>
                    <span onClick={() => this.handleDeleteConfig(record.key)}>
                        <a href="javascript:;">删除</a>
                    </span>
                </div>
            ),
        }];

        return (
            <div>
                <div>
                    <span className="sub-title">Config</span>
                </div>
                <div className="filter-bar">
                    <div className="search-input">
                        <Search
                            placeholder="input search text"
                            onSearch={value => console.log(value)}
                            style={{ width: 200 }}
                        />
                    </div>
                    <Button type="primary" onClick={() => this.handleCreateConfig()}>Create</Button>
                </div>
                <Table columns={columns} dataSource={this.state.data} loading={this.state.loading} />
                <ConfigModal title={this.state.modalTitle} visible={this.state.modalVisible} onCancel={() => this.handleCancel()} modalData={this.state.modalData} loadData={()=>this.loadData()}/>
            </div>
        );
    }
}

export default ConfigList;
