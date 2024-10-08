## User stories

A user story is implemented as well as it is communicated.
If the context and the goals are made clear, it will be easier for everyone to implement it, test it, refer to it…

---

Quick links: [Summary](#summary) | [Description](#description) | [Template](#template) | [Example](#example) | [Resources](#resources)

---

### Summary

A user story should typically have a summary structured this way:

1. **As a** [user concerned by the story]
1. **I want** [goal of the story]
1. **so that** [reason for the story]

The “so that” part is optional if more details are provided in the description.

This can then become “As a user managing my properties, I want notifications when adding or removing images.”

You can read about some reasons for this structure in this [nicely put article][1].

### Description

We’re using the following template to create user stories. 

Since we have mentioned the type of user, the user story can refer to it with “I”.
This is useful for **consistency** and to **avoid repetition** in the Acceptance criteria.
It’s also good to practice a little **empathy**.

For example:

```markdown
1. I click on the “submit” button.
1. A modal window appears if I don’t have enough credits.
1. The modal window contains the following:
  1. […]
```

The template uses [markdown][2].
You should get familiar with it if you’re not already, **it’s awesome!**

### Template

```markdown
[
The user story should have a reason to exist: what do I need as the user described in the summary?
This part details any detail that could not be passed by the summary.
]


### Acceptance Criteria

1. [If I do A.]
1. [B should happen.]

[
Also, here are a few points that need to be addressed:

1. Constraint 1;
1. Constraint 2;
1. Constraint 3.
]

### Definition of Done

### Resources:

* Mockups: [Here goes a URL to or the name of the mockup(s) in inVision];
* Testing URL: [Here goes a URL to the testing branch or IP];
* Staging URL: [Here goes a URL to the feature on staging];


### Notes

[Some complementary notes if necessary:]

* > Here goes a quote from an email
* Here goes whatever useful information can exist…
```

### Example
```markdown

Title: Implement kube-egress-gateway for Fixed Source IP Egress in AKS

As a DevOps engineer,
I want to implement kube-egress-gateway in our AKS cluster,
So that we can efficiently manage and secure outbound traffic from pods with fixed source IPs.

Acceptance Criteria:

1. kube-egress-gateway components are successfully deployed in the AKS cluster.
2. A dedicated node pool is set up as the gateway VMSS for egress traffic.
3. StaticGatewayConfiguration Custom Resource is created and configured correctly.
4. Public IP prefix is provisioned and associated with the gateway VMSS.
5. Network policies are implemented to control traffic flow between pods and the egress gateway.
6. RBAC is configured to manage access to StaticGatewayConfiguration resources.
7. All communication between pods and the gateway is encrypted.
8. Logging and monitoring are set up for the kube-egress-gateway components.
9. Pod annotations for using the egress gateway are documented and tested.
10. Egress filtering is implemented on the gateway to control outbound connections.
11. A disaster recovery plan for the egress gateway is documented and tested.
12. Security hardening measures are implemented as per the recommendations.
13. Performance impact of routing traffic through the egress gateway is measured and documented.
14. User documentation is created for developers on how to use the egress gateway for their pods.
15. An automated test suite is developed to verify the functionality and security of the egress gateway setup.

Definition of Done:

1. All acceptance criteria are met and verified.
2. Code review is completed and approved.
3. All automated tests pass.
4. Documentation is updated and reviewed.
5. The solution is deployed to a staging environment and tested.
6. Performance benchmarks are within acceptable ranges.
7. Security team has audited and approved the implementation.
8. Runbook for operations team is created and reviewed.

Story Points: 13 (This is a complex task involving multiple components and requiring careful testing and security considerations)

Notes:
- Coordinate with the security team for approval of the implementation details.
- Consider breaking this story into smaller, more manageable tasks during sprint planning.
- Ensure compliance with any relevant regulatory requirements for network traffic management.







