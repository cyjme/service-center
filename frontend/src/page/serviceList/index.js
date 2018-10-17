import React, { Component } from 'react';
import './index.css';
import { message, Input, Radio, Table } from 'antd';
import request from '../../utils/request';

const Search = Input.Search;
class ServiceList extends Component {
    constructor() {
        super()
        this.state = {
            loading: true,
            services: []
        }
    }
    componentDidMount() {
        this.loadData()
    }
    async loadData() {
        this.setState({
            loading: true
        })
        const res = await request.get("/service")
        if (res.status !== 200) {
            message.error("get service error")
            return
        }
        let services = [];
        for (const key in res.data) {
            services.push({
                key: Math.random(),
                name: key,
                nodes: res.data[key]
            })
        }

        this.setState({
            services: services
        })
    }
    render() {
        const columns = [{
            title: 'name',
            dataIndex: 'name',
            key: 'name',
            render: text => <a href="javascript:;">{text}</a>,
        }, {
            title: 'nodeNumber',
            dataIndex: 'nodes',
            key: 'nodeNumber',
            render: nodes => nodes.length
        }, {
            title: 'Nodes',
            dataIndex: 'nodes',
            key: 'nodes',
            render: nodes => {
                return (
                    <div>
                        {nodes.map(item => {
                            return (
                                <div className="node" key={item.key}>
                                    <p>
                                        key: {item.key}
                                    </p>
                                    <p>
                                        value: {item.url}
                                    </p>
                                </div>
                            )
                        })}
                    </div>
                )
            },
        }, {
            title: 'Action',
            key: 'action',
            render: (text, record) => (
                <span>
                    <a href="javascript:;">详情</a>
                </span>
            ),
        }];

        return (
            <div>
                <div>
                    <span className="sub-title">Services</span>
                </div>
                <div className="filter-bar">
                    <div className="radio-group">
                        <Radio.Group size="default" value={"all"} onChange={this.handleSizeChange}>
                            <Radio.Button value="all">All</Radio.Button>
                            <Radio.Button value="passing">Passing</Radio.Button>
                            <Radio.Button value="warning">Warning</Radio.Button>
                            <Radio.Button value="critical">Critical</Radio.Button>
                        </Radio.Group>
                    </div>
                    <div className="search-input">
                        <Search
                            placeholder="input search text"
                            onSearch={value => console.log(value)}
                            style={{ width: 200 }}
                        />
                    </div>
                </div>
                <Table columns={columns} dataSource={this.state.services} />
            </div>
        );
    }
}

export default ServiceList;
