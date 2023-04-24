import React from 'react';
import './index.css';
import Host from '../host';

class Group extends React.Component {
    constructor (props) {
        super(props);
        this.state = {
            name: this.props.groupName,
            description: this.props.groupDescription,
            hosts: null
        }
    }


    componentDidMount() {
        this.loadData(); // Fetch data immediately
        this.interval = setInterval(() => this.loadData(), 5000); // Fetch data every 8 seconds
    }
    
    componentWillUnmount() {
        clearInterval(this.interval); // Clear the interval when the component unmounts
    }

    loadData() {
        fetch('/getGroupAvailable/'+this.props.groupName)
          .then(response => response.json())
          .then(data => {
            this.setState({ hosts: data, error: null });
          })
          .catch(error => {
            this.setState({ error: error });
          });
    }

    createHostsElements(){
        
        let hostsElements = this.state.hosts.map((host) => {
            return <Host key={host.ID} Hostname={host.Hostname} ipAddress={host.HostIP} alive={host.IsAlive}></Host>
        })
        
        return hostsElements
    }

    render() {

        const { description, name, hosts, error } = this.state;

        if (error) {
          return <div>Error: {error.message}</div>;
        }
    
        if (!hosts) {
          return <div>Loading...</div>;
        }
                
        return (
            <div className='group' style={{
                backgroundColor: 'rgba(39,41,83)'
            }}>
                <div className='groupHeader'>
                    {name}
                    <br></br>
                    {description}
                </div>

                {this.createHostsElements()}
            </div>
        );
    }
}

export default Group