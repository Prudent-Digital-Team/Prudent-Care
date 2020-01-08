<template>
  <div class="faq">
    <div class="pg-title is-blue-bold is-size-3">Frequently Asked Questions (FAQ)</div>
    <div class="container">
      <!-- <p class="is-blue-bold is-size-5">General</p> -->
      <div class="faq-list">
        <b-collapse class="card faq-card" :open="isOpen" v-for="faq in FaqList" :key="faq.id">
          <div slot="trigger" class="card-header" role="button" aria-controls="contentIdForA11y3">
            <p class="card-header-title is-blue-bold is-size-6">{{faq.title}}</p>
            <a class="card-header-icon">
              <b-icon :icon="isOpen ? 'menu-down' : 'menu-up'"></b-icon>
            </a>
          </div>
          <div class="card-content">
            <div class="content is-blue is-glight is-size-5">{{faq.content}}</div>
          </div>
        </b-collapse>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  head() {
    return {
      title: "Faqs"
    };
  },
  async asyncData({ store }) {
    store.commit("setbanner", "faq.jpg");
  },
  data() {
    return {
      isOpen: false,
      FaqList: [
        {
          title: "How do I pay for my care? ",
          content: `We will send you a bill once a fortnight. If you wish to pay by standing order (to save you having to remember to write out a cheque) just let us know. `
        },
        {
          title: "Knowing how much to pay",
          content: `Every two weeks we will send you an invoice. This is a request for payment for the care and support you have received.  We can send your invoice by post or email. Let us know which you would prefer by contacting a member of the Prudent Domiciliary Care team`
        },
        {
          title: "When will payment be due?",
          content:
            "Prudent Domiciliary Care request payment one week following the date of invoice. "
        },
        {
          title: "Set up a standing order",
          content: `If you choose to pay by standing order, we will take a specified amount from your nominated account every calendar month.
This may not always fully cover the cost of your care and support. For example, if you have received extra care because of a bank holiday or extra visits for any other reason. Where this is the case our finance team will send you a statement to tell you about the outstanding balance on your account.
If you would like to set up a standing order, please contact us so we can send you the documents you need for your bank. 
`
        },
        {
          title: "Set up a direct debit",
          content: `If you choose to pay by direct debit we will take a specified amount from your nominated account following the issue of your invoice.
Direct debits offer more flexibility than standing orders. You can choose to pay a regular amount as you would with a standing order. Alternatively, you can choose to pay different amounts each month to pay for the actual cost of your care and support for that period.
If you would like to set up a direct debit, please contact us so we can send you the documents you need for your bank.  
`
        },
        {
          title: "Make a payment by debit or credit card over the phone",
          content: `Prudent Domiciliary Care are unable to take payment over the telephone.`
        },
        {
          title: "Transfer the payment electronically",
          content: `You can pay our invoice by electronic transfer/internet banking. Please contact Prudent Domiciliary Care to obtain our bank details (account name, sort code and account number) you will need these to transfer the payment.  `
        },
        {
          title: "Send us a cheque",
          content: `You can pay our invoice by cheque`
        },
        {
          title:
            "How do I arrange a care service with Prudent Domiciliary Care? ",
          content:
            "Contact Prudent Domiciliary Care and make an appointment for us to visit you for a chat about how Prudent Domiciliary Care can work for you. We are happy to visit you at home to discuss ways we can help and then, once you have decided, you simply leave the rest to us!"
        },
        {
          title:
            "What if I don't like my care package or want to make changes?",
          content:
            "Our experience tells us that people's care needs very often change. We therefore understand that care may need reviewing from time to time. This is why we carry out reviews on a regular basis and why we are pleased to hear from members of the immediate family about effective methods for delivering the care service for their relative. "
        },
        {
          title: "What tasks will my care assistant carry out?",
          content: `We agree a ‘Care Plan’ which details exactly what you have instructed us to do. This then becomes the work list for each day. Click here for details of the care services we can provide
Care plans frequently evolve and we check regularly to ensure that the care we are providing is appropriate. 
`
        },
        {
          title: "Is there anything that my care assistant is not able to do?",
          content: `Once we have established exactly what you want on each visit, we will work to that list. We are not able to give injections or change sterile dressings. Also, we cannot carry out tasks which are likely to cause risk to either the person we provide care to or to the care assistant. Tasks will be agreed with you in advance. Any changes to the plan can be discussed with Prudent Domiciliary Care team.  `
        },
        {
          title: "Will I have the same care assistant each day?",
          content: `This depends on how many care visits you have each week. We very rarely have just one care assistant assigned to a particular customer. This allows for periods of time where your care assistant may be away on holiday, however we do try to keep the number of your care assistants to an absolute minimum. `
        },
        {
          title: "What about my personal security?",
          content: `Every member of the Prudent Domiciliary Care team has been through a rigorous vetting procedure before they start working with us. In addition, all staff will always be in uniform and carry a personalised identification card which shows their name, photograph and signature. `
        },
        {
          title:
            "How can I be sure that my care assistant will stay the full period of the care visit and carry out all the tasks? ",
          content: `We take a considered and comprehensive approach to ensuring that your care service is exactly what is required. After every visit the care assistant will write the details of activities undertaken during the visit in the care record (which is retained by the customer). In addition, the customer is asked to record their satisfaction by signing the care assistant's timesheet, which is returned to the Prudent Domiciliary Care office each week for checking. `
        },

        {
          title:
            "How much do Prudent Domiciliary Care services cost and can I get any help with the fees? ",
          content: `We will put together a detailed plan for each customer which, as well as setting out the care to be delivered, also shows the weekly costs. Some people will be entitled to assistance with the cost of their care either from their local Social Services office or via other Government grants.`
        },
        {
          title: "What happens when the office is closed?",
          content: `Our offices are normally open from 9am to 5pm Monday to Friday. Aside from this, we operate an out-of-hours emergency service which is manned from 7am to 10pm. The number for your out of hours service can be provided for your reassurance, with someone at the end of the telephone line to assist with any emergencies about your care. `
        }
      ]
    };
  }
};
</script>

<style lang="scss" scoped></style>
