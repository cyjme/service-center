import React from "react"
import { Form, Modal, Input ,message} from "antd";
import request from '../../../../utils/request';

const { TextArea } = Input;
const FormItem = Form.Item;

class ConfigModal extends React.Component {
    async updateConfig(key, value) {
        const res = await request.put("/config", { key: key, value: value })
        if (res.status != 200) {
            message.error("create fail")
            return
        }
        message.success("create success")
        this.props.loadData()
        this.props.onCancel()
    }

    handleSubmit = (e) => {
        e.preventDefault();
        this.props.form.validateFields((err, values) => {
            if (!err) {
                console.log('Received values of form: ', values);
                this.updateConfig(values.key,values.value)
            }
        })
    }
    render() {
        const { getFieldDecorator } = this.props.form;
        return (
            <div>
                <Modal
                    title={this.props.title}
                    visible={this.props.visible}
                    onOk={this.handleSubmit}
                    onCancel={this.props.onCancel}
                    destroyOnClose={true}
                >
                    <Form layout={'vertical'}>
                        <FormItem
                            label="key"
                        >
                            {getFieldDecorator('key', {
                                rules: [{ required: true, message: 'Please input key !' }],
                                initialValue: this.props.modalData.key
                            })(
                                <Input placeholder="please input key" disabled={this.props.modalType=="edit"?true:false}/>
                            )}
                        </FormItem>

                        <FormItem
                            label="value"
                        >
                            {getFieldDecorator('value', {
                                initialValue: this.props.modalData.value
                            })(
                                <TextArea rows={6} placeholder="please input value" />
                            )}
                        </FormItem>
                    </Form>
                </Modal>
            </div>
        )
    }
}

export default Form.create()(ConfigModal);