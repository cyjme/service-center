import React from 'react';
import './index.css'
import { withRouter } from 'react-router'

const menu = [
    {
        key: "service",
        name: "Service",
        path: "/",
    },
    {
        key: "config",
        name: "Config",
        path: "/config",
    },
]
class MainLayout extends React.Component {
    constructor() {
        super()
        this.state = {
            selected: "service"
        }
    }

    handleSelectMenu(item) {
        this.setState({
            selected: item.key
        })
        this.props.history.push(item.path)
    }

    renderHeader() {
        return (
            <header className="header">
                <span className="title">Welcome</span>
                {
                    menu.map(item => {
                        return (
                            <span key={item.key} className={this.state.selected === item.key ? "menu-item active" : "menu-item"} onClick={() => this.handleSelectMenu(item)}>
                                {item.name}
                            </span>
                        )
                    })
                }
            </header>
        )
    }

    render() {
        return (
            <div className="App" location={this.props.location}>
                {this.renderHeader()}
                <div className="main-container">
                    {this.props.children}
                </div>
            </div>
        )
    }
}
export default withRouter(MainLayout);